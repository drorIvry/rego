import Link from "next/link"
import MaxWidthWrapper from "./maxWidthWrapper"
import { Button, buttonVariants } from "./ui/button"
import LoginLink from "./loginLink"

const Navbar = () => {
  return (
    <nav className="sticky h-14 inset-x-0 top-0 z-30 w-full border-b border-gray-900 bg-black/75 backdrop-blur-lg transition-all">
      <MaxWidthWrapper>
        <div className="flex h-14 items-center justify-between border-b">
          <Link 
            href="/" 
            className="flex z-40 font-semibold"
          >
            <span>Rego</span>
          </Link>

          {/* TODO: add mobile bavbar */}

          <div className="hidden items-center space-x-4 sm:flex">
            <>
              <Link 
                className={buttonVariants({
                  variant: "ghost",
                  size: "sm",
                })}
                href="/pricing"
              >
                Pricing
              </Link>
              <LoginLink/>
            </>
          </div>

        </div>
      </MaxWidthWrapper>
    </nav>
  )
}


export default Navbar
