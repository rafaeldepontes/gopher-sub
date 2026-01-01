"use client";

import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

const SUFIX = "/api/v1";

type StoredUser = { id?: string; token?: string } | null;

const SubButton = () => {
    const router = useRouter();
    const [user, setUser] = useState<StoredUser>(null);
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        try {
            const raw = localStorage.getItem("user");
            if (!raw) {
                router.push("/login");
                return;
            }
            setUser(JSON.parse(raw));
        } catch (err) {
            console.error("Failed to read user from localStorage", err);
            router.push("/login");
        }
    }, [router]);

    if (!user) return null;

    const handleSubmit = async () => {
        setLoading(true);
        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_URL ?? ""}${SUFIX}/subscribe/${user.id}`,
                {
                    method: "POST",
                    headers: {
                        Authorization: user.token ?? "",
                    },
                }
            );

            if (!res.ok) {
                const text = await res.text().catch(() => "");
                throw new Error(`Subscribe failed: ${res.status} ${text}`);
            }
            window.location.replace("http://localhost:8025");
        } catch (err) {
            console.error(err);
            alert("Subscription failed. Check console for details.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <button
            onClick={handleSubmit}
            disabled={loading}
            className="w-full rounded-md bg-white text-black py-2 text-sm font-medium hover:opacity-90 transition disabled:opacity-50"
        >
            {loading ? "Subscribing..." : "Subscribe"}
        </button>
    );
};

export default SubButton;
