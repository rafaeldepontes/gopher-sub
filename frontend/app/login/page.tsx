import FormsButton from "@/components/FormsButton"

const inputClass =
    "w-full rounded-md bg-[var(--surface)] border border-[var(--border)] px-3 py-2 text-sm text-[var(--text)] placeholder-[var(--muted)] focus:outline-none focus:border-white focus:ring-1 focus:ring-white"

const LoginPage = () => {
    return (
        <div className="min-h-screen flex items-center justify-center">
            <div className="w-full max-w-sm rounded-xl border border-[var(--border)] bg-[var(--surface)] p-6 flex flex-col gap-4">
                <h1 className="text-lg font-semibold">Login</h1>

                <input className={inputClass} placeholder="Email" />
                <input className={inputClass} placeholder="Password" type="password" />

                <FormsButton isLogin />
            </div>
        </div>
    )
}

export default LoginPage
