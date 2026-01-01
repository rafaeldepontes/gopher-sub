"use client";

import SubButton from "@/components/SubButton";

const HomePage = () => {
    return (
        <div className="min-h-screen flex items-center justify-center">
            <div className="w-full max-w-sm rounded-xl border border-[var(--border)] bg-[var(--surface)] p-6 flex flex-col gap-4">
                <h1 className="text-lg font-semibold">Hi random user!</h1>

                <p className="text-sm text-[var(--muted)]">
                    Click here to receive news about our product:
                </p>

                <SubButton />
            </div>
        </div>
    );
};

export default HomePage;
