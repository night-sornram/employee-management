'use client'

import { useRouter } from "next/navigation";

export default function LoginPage () {

    const router = useRouter();

    const handleLogin= async (e: React.FormEvent<HTMLFormElement>) => {
        alert("login!");
        router.push('/attendance');
        router.refresh();
    }

    return (
      <div className="flex flex-row h-full w-[85vw]">
        <div className="flex flex-col w-full h-full">
            <div className="px-[10%] py-[10%]">
                <h1 className="text-2xl font-bold mb-[5%]">Login</h1>
                <form onSubmit={handleLogin}>
                    <div className="mb-[5%]">
                        <h2>ID</h2>
                        <input type="text" placeholder="Your ID" className="bg-slate-100 placeholder-slate-400 rounded px-3 py-2 w-[40%]
                        border border-slate-300 shadow-sm focus:outline-none focus:border-slate-400 invalid:border-red-600 invalid:text-red-600 focus:invalid:border-red-700"/>
                    </div>
                    <div className="mb-[5%]">
                        <h2>Password</h2>
                        <input type="password" placeholder="Password" className="bg-slate-100 placeholder-slate-400 rounded px-3 py-2 w-[40%]
                        border border-slate-300 shadow-sm focus:outline-none focus:border-slate-400 invalid:border-red-600 invalid:text-red-600 focus:invalid:border-red-700"/>
                    </div>
                    <div className="flex flex-col items-center w-[40%]">
                        <button type="submit" className="custom-btn-dark hover:custom-btn-dark-hover">
                            Login
                        </button>
                    </div>  
                </form>
            </div>
        </div>
      </div>  
    );
}