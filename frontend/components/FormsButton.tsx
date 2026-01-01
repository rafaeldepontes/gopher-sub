"use client"
import Link from "next/link"

const FormsButton = ({
    isLogin,
    handleSubmit
}: {
    isLogin: boolean,
    handleSubmit: () => Promise<void>
}) => {
    return (
        <div className="flex flex-col gap-3">
            <button className="w-full rounded-md bg-white text-black py-2 text-sm font-medium hover:opacity-90 transition" onClick={() => handleSubmit()}>
                {isLogin ? "Login" : "Register"}
            </button>

            <p className="text-xs text-[var(--muted)] text-center">
                {isLogin ? (
                    <>
                        Don&apos;t have an account?{" "}
                        <Link href="/register" className="text-white underline">
                            Sign up
                        </Link>
                    </>
                ) : (
                    <>
                        Already have an account?{" "}
                        <Link href="/login" className="text-white underline">
                            Login
                        </Link>
                    </>
                )}
            </p>
        </div>
    )
}

export default FormsButton
