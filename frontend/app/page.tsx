import { redirect, RedirectType } from "next/navigation"

const Root = () => {
    redirect('/login', RedirectType.push)
}

export default Root