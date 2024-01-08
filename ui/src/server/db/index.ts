import { neon } from "@neondatabase/serverless";
import { drizzle } from "drizzle-orm/neon-http";

import { env } from "~/env";
import * as schema from "./schema";

// See https://neon.tech/docs/serverless/serverless-driver
// for more information
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
const sql = neon(env.DATABASE_URL);
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
export const db = drizzle(sql, schema);
