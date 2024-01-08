import { neon } from "@neondatabase/serverless";
import { drizzle } from "drizzle-orm/neon-http";

import { env } from "~/env";
import * as schema from "./schema";

const posts = await sql("SELECT * FROM posts");

// See https://neon.tech/docs/serverless/serverless-driver
// for more information
const sql = neon(env.DATABASE_URL);
export const db = drizzle(sql, schema);
