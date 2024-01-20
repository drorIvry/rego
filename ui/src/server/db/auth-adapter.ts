import type {
  AdapterAccount,
  AdapterSession,
  AdapterUser,
} from "next-auth/adapters";
import {
  users as usersTable,
  sessions as sessionsTable,
  accounts as AccountsTable,
  verificationTokens as verificationTokens,
} from "./schema";
import { and, eq } from "drizzle-orm";
import { createId } from "@paralleldrive/cuid2";
import type { VerificationToken } from "next-auth/adapters";
import type { NeonDatabase } from "drizzle-orm/neon-serverless";

export function sqlDrizzleAdapter(client: NeonDatabase) {
  return {
    async createUser(data: Omit<AdapterUser, "id">) {
      const id = createId();

      await client.insert(usersTable).values({ ...data, id });

      return await client
        .select()
        .from(usersTable)
        .where(eq(usersTable.id, id))
        .then((res) => res[0]);
    },
    async getUser(data: string) {
      const thing =
        (await client
          .select()
          .from(usersTable)
          .where(eq(usersTable.id, data))
          .then((res) => res[0])) ?? null;
      return thing;
    },
    async getUserByEmail(data: string) {
      const user =
        (await client
          .select()
          .from(usersTable)
          .where(eq(usersTable.email, data))
          .then((res) => res[0])) ?? null;

      return user;
    },
    async createSession(data: {
      sessionToken: string;
      userId: string;
      expires: Date;
    }) {
      await client.insert(sessionsTable).values({ ...data, id: createId() });

      return await client
        .select()
        .from(sessionsTable)
        .where(eq(sessionsTable.sessionToken, data.sessionToken))
        .then((res) => res[0]);
    },
    async getSessionAndUser(data: string) {
      const rows = await client
        .select({
          user: usersTable,
          session: {
            id: sessionsTable.id,
            userId: sessionsTable.userId,
            sessionToken: sessionsTable.sessionToken,
            expires: sessionsTable.expires,
          },
        })
        .from(sessionsTable)
        .innerJoin(usersTable, eq(usersTable.id, sessionsTable.userId))
        .where(eq(sessionsTable.sessionToken, data))
        .limit(1);

      if (!rows) return null;
      const row = rows[0];
      if (!row) return null;

      const { user, session } = row;
      return {
        user,
        session: {
          id: session.id,
          userId: session.userId,
          sessionToken: session.sessionToken,
          expires: session.expires,
        },
      };
    },
    async updateUser(data: Partial<AdapterUser> & Pick<AdapterUser, "id">) {
      if (!data.id) {
        throw new Error("No user id.");
      }

      await client
        .update(usersTable)
        .set(data)
        .where(eq(usersTable.id, data.id));

      return await client
        .select()
        .from(usersTable)
        .where(eq(usersTable.id, data.id))
        .then((res) => res[0]);
    },
    async updateSession(
      data: Partial<AdapterSession> & Pick<AdapterSession, "sessionToken">,
    ) {
      await client
        .update(sessionsTable)
        .set(data)
        .where(eq(sessionsTable.sessionToken, data.sessionToken));

      return await client
        .select()
        .from(sessionsTable)
        .where(eq(sessionsTable.sessionToken, data.sessionToken))
        .then((res) => res[0]);
    },
    async linkAccount(rawAccount: AdapterAccount) {
      await client
        .insert(AccountsTable)
        .values({ ...rawAccount, id: createId() });
    },
    async getUserByAccount(
      account: Pick<AdapterAccount, "providerAccountId" | "provider">,
    ) {
      const dbAccount =
        (await client
          .select()
          .from(AccountsTable)
          .where(
            and(
              eq(AccountsTable.providerAccountId, account.providerAccountId),
              eq(AccountsTable.provider, account.provider),
            ),
          )
          .leftJoin(usersTable, eq(AccountsTable.userId, usersTable.id))
          .then((res) => res[0])) ?? null;

      console.log(">>>>>>", dbAccount);
      if (!dbAccount) {
        return null;
      }

      return dbAccount.user;
    },
    async deleteSession(sessionToken: string) {
      const session =
        (await client
          .select()
          .from(sessionsTable)
          .where(eq(sessionsTable.sessionToken, sessionToken))
          .then((res) => res[0])) ?? null;

      await client
        .delete(sessionsTable)
        .where(eq(sessionsTable.sessionToken, sessionToken));

      return session;
    },
    async createVerificationToken(token: VerificationToken) {
      await client.insert(verificationTokens).values(token);

      return await client
        .select()
        .from(verificationTokens)
        .where(eq(verificationTokens.identifier, token.identifier))
        .then((res) => res[0]);
    },
    async useVerificationToken(token: { identifier: string; token: string }) {
      try {
        const deletedToken =
          (await client
            .select()
            .from(verificationTokens)
            .where(
              and(
                eq(verificationTokens.identifier, token.identifier),
                eq(verificationTokens.token, token.token),
              ),
            )
            .then((res) => res[0])) ?? null;

        await client
          .delete(verificationTokens)
          .where(
            and(
              eq(verificationTokens.identifier, token.identifier),
              eq(verificationTokens.token, token.token),
            ),
          );

        return deletedToken;
      } catch (err) {
        throw new Error("No verification token found.");
      }
    },
    async deleteUser(id: string) {
      const user = await client
        .select()
        .from(usersTable)
        .where(eq(usersTable.id, id))
        .then((res) => res[0] ?? null);

      await client.delete(usersTable).where(eq(usersTable.id, id));

      return user;
    },
    async unlinkAccount(
      account: Pick<AdapterAccount, "providerAccountId" | "provider">,
    ) {
      await client
        .delete(AccountsTable)
        .where(
          and(
            eq(AccountsTable.providerAccountId, account.providerAccountId),
            eq(AccountsTable.provider, account.provider),
          ),
        );

      return undefined;
    },
  };
}
