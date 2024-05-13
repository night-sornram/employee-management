import Link from "next/link";

export default function NavBar () {
    return (
        <div className="bg-p-gray w-[15%] h-screen flex flex-col justify-start items-center">
            <Link href='/attendance' className="bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center">
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Attendance</h1>
                </div>
            </Link>
            <Link href='/attendance' className="bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center">
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Leave Request</h1>
                </div>
            </Link>
            <Link href='/attendance' className="bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center">
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Attendance History</h1>
                </div>
            </Link>
            <Link href='/attendance' className="bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center">
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Leave History</h1>
                </div>
            </Link>
            <Link href='/attendance' className="bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center">
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Leave Management</h1>
                </div>
            </Link>
        </div>
    );
}