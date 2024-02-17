import Link from "next/link";

import MaxWidthWrapper from "~/components/maxWidthWrapper";
import { getServerAuthSession } from "~/server/auth";
import { api } from "~/trpc/server";
import { ArrowRight } from "lucide-react";
import { buttonVariants } from "~/components/ui/button";
import Image from "next/image"

export default async function Home() {
  const hello = await api.post.hello.query({ text: "from tRPC" });
  const session = await getServerAuthSession();

  return (
    <>
      <MaxWidthWrapper className="mb-12 mt-28 sm:mt-40 flex flex-col items-center justify-center text-center">
        <div className="mx-auto mb-4 flex max-w-fit items-center justify-center space-x-2 overflow-hidden rounded-full border border-gray-200 bg-white px-7 py-2 shadow-md backdrop-blur transition-all hover:border-gray-300 hover:bg-white/50">
          <p className="text-sm font-semibold text-gray-700">
            Rego test
          </p>
        </div>
        <h1 className="max-w-4xl text-5xl font-bold md:text-6xl lg:text-7xl">
          Rego
        </h1>
        <p className="mt-5 max-w-prose text-zinc-700 sm:text-lg">
          Rego
        </p>
        <Link className={buttonVariants({
          size: "lg",
          className: "mt-5",
        })} href="/dashboard" target='_blank'>
          Get started <ArrowRight className="ml-2 h-5 w-5"/>
        </Link>
      </MaxWidthWrapper>

      <div>
        <div className="relative isolate">
          <div 
            aria-hidden="true" 
            className="pointer-events-none absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
          >
            <div style={{
              clipPath: "polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
            }} className="relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]" 
            />
          </div>

          <div>
            <div className="mx-auto max-w-xl px-6 lg:px-8">
              <div className="mx-auto mt-16 flow-root sm:mt-24">
                  <Image 
                    src="/rego.png"
                    alt="Rego Logo"
                    width={529}
                    height={471}
                    quality={100}
                  />
              </div>
            </div>
          </div>

          <div 
            aria-hidden="true" 
            className="pointer-events-none absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
          >
            <div style={{
              clipPath: "polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
            }} className="relative left-[calc(50%-13rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-36em)] sm:w-[72.1875rem]" 
            />
          </div>
        </div>
      </div>

      {/* Feature section */}
      <div className='mx-auto mb-32 mt-16 max-w-5xl sm:mt-16'>
        <div className="mb-12 px-6 lg:px-8">
          <div className="mx-auto max-w-2xl sm:text-center">
            <h2 className="mt-2 font-bold text-4xl text-gray-200 sm:text-5xl">
              Execute your docker at edge in seconds!
            </h2>
            <p className="mt-4 text-lg text-gray-200">
              Execute your job once or periodically at edge using rego
            </p>
          </div>
        </div>

        {/* Steps */}
        <ol className="my-8 space-y-4 pt-8 md:flex md:space-x-12 md:space-y-0">
          <li className="md:flex-1">
            <div className="flex flex-col space-y-2 border-l-4 border-zinc-700 py-2 pl-4 md:border-l-0 md:border-t-2 md:pb-0 md:pl-0 md:pt-4">
              <span className="text-sm font-medium text-blue-600">Step 1</span>
              <span className="text-xl font-semibold">Sign up for an account</span>
              <span className="mt-2 text-zinc-700">Try out</span>
            </div>
          </li>
          <li className="md:flex-1">
            <div className="flex flex-col space-y-2 border-l-4 border-zinc-700 py-2 pl-4 md:border-l-0 md:border-t-2 md:pb-0 md:pl-0 md:pt-4">
              <span className="text-sm font-medium text-blue-600">Step 2</span>
              <span className="text-xl font-semibold">Create an API key</span>
              <span className="mt-2 text-zinc-700">
                You can do so {' '}
                <Link 
                  href="/api-keys"
                  className="text-blue-700 underline underline-offset-2"
                >
                  here
                </Link>
              </span>
            </div>
          </li>
          <li className="md:flex-1">
            <div className="flex flex-col space-y-2 border-l-4 border-zinc-700 py-2 pl-4 md:border-l-0 md:border-t-2 md:pb-0 md:pl-0 md:pt-4">
              <span className="text-sm font-medium text-blue-600">Step 3</span>
              <span className="text-xl font-semibold">Create a task for your docker</span>
              <span className="mt-2 text-zinc-700">You can configure your task to run once or periodically based on your needs</span>
            </div>
          </li>
        </ol>
      </div>
    </>
  );
}
