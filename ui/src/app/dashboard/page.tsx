import { redirect } from "next/navigation"
import Dashboard from "~/components/dashboard"
import { getServerAuthSession } from "~/server/auth"

const Page = async () => {
    const session = await getServerAuthSession()
    const user = session?.user

    if (!user || !user.id) {
        redirect("/api/auth/signin")
    }
        
    return <Dashboard />
}

export default Page
