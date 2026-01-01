"use client"
import { useState } from "react"
import FormsButton from "@/components/FormsButton"
import { redirect, RedirectType } from "next/navigation"

const inputClass =
    "w-full rounded-md bg-[var(--surface)] border border-[var(--border)] px-3 py-2 text-sm text-[var(--text)] placeholder-[var(--muted)] focus:outline-none focus:border-white focus:ring-1 focus:ring-white"

const SUFIX = "/api/v1"

const LoginPage = () => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    async function handleSubmit() {
        const res = await fetch(`${process.env.NEXT_PUBLIC_URL}${SUFIX}/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, password }),
        })

        if (!res.ok) {
            throw new Error("Login failed")
        }

        const data = await res.json()
        localStorage.setItem("user", JSON.stringify(data))

        redirect("/home", RedirectType.push)
    }

    return (
        <div className="min-h-screen flex items-center justify-center">
            <div className="w-full max-w-sm rounded-xl border border-[var(--border)] bg-[var(--surface)] p-6 flex flex-col gap-4">
                <h1 className="text-lg font-semibold">Login</h1>

                <input
                    className={inputClass}
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="Email"
                />
                <input
                    className={inputClass}
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="Password"
                    type="password"
                />

                <FormsButton isLogin handleSubmit={handleSubmit} />
            </div>
        </div>
    )
}

export default LoginPage
