"use client"
import FormsButton from "@/components/FormsButton"
import { redirect, RedirectType } from "next/navigation"
import { useState } from "react"

const inputClass =
    "w-full rounded-md bg-[var(--surface)] border border-[var(--border)] px-3 py-2 text-sm text-[var(--text)] placeholder-[var(--muted)] focus:outline-none focus:border-white focus:ring-1 focus:ring-white"

const SUFIX = "/api/v1"

const RegisterPage = () => {
    const [email, setEmail] = useState<string>("")
    const [password, setPassword] = useState<string>("")
    const [age, setAge] = useState<number>(0)

    const handleSubmit = async () => {
        if (email == "" || password == "" || age < 13) {
            return
        }

        await fetch(`${process.env.NEXT_PUBLIC_URL}${SUFIX}/register`, {
            method: "POST",
            headers: {
                "Content-Type": "applcation/json",
            },
            body: JSON.stringify({ email, password, age }),
        })

        redirect("/login", RedirectType.push)
    }

    return (
        <div className="min-h-screen flex items-center justify-center">
            <div className="w-full max-w-sm rounded-xl border border-[var(--border)] bg-[var(--surface)] p-6 flex flex-col gap-4">
                <h1 className="text-lg font-semibold">Register</h1>
                <input
                    className={inputClass}
                    placeholder="Email"
                    value={email}
                    onChange={(event) => setEmail(event.target.value)}
                    type="text"
                />
                <input
                    className={inputClass}
                    placeholder="Password"
                    value={password}
                    onChange={(event) => setPassword(event.target.value)}
                    type="password"
                />
                <input
                    className={inputClass}
                    placeholder="Age"
                    value={age}
                    onChange={(event) => setAge(parseInt(event.target.value))}
                    type="number"
                />
                <FormsButton isLogin={false} handleSubmit={handleSubmit} />
            </div>
        </div>
    )
}

export default RegisterPage