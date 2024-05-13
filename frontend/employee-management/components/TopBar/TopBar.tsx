import Link from "next/link";

export default function TopBar () {
    return (
        <div className="bg-p-darkgray h-[6vh] w-full flex flex-row justify-end pr-[1%]">
            <div className="flex flex-col justify-center px-6">
                <h1 className="text-white text-lg">Prasut</h1>
            </div>
            <Link href="/login" className="flex flex-col justify-center px-6 hover:bg-zinc-800">
                <div>
                    <h1 className="text-white text-lg">Logout</h1>
                </div>
            </Link>
        </div>
    );
}