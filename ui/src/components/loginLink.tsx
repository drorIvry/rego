import Link from "next/link";
import { buttonVariants } from "./ui/button";
import { getServerAuthSession } from "~/server/auth";

const LoginLink = async () => {
  const session = await getServerAuthSession();

  return (
    <>
      {/* <p className="text-center text-2xl text-white">
        {session && <span>Logged in as {session.user?.name}</span>}
      </p> */}

      <Link
        href={session ? "/api/auth/signout" : "/api/auth/signin"}
        className={buttonVariants({
          className: "bg-gray",
          size: "sm",
        })}
      >
        {session ? "Sign out" : "Sign in"}
      </Link>
    </>
  );
};

export default LoginLink;
