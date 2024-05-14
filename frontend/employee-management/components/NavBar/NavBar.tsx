'use client'
import { usePathname, useRouter } from "next/navigation";

export default function NavBar () {

    const router = useRouter();
    const path = usePathname();

    return (
        <div className="bg-p-gray w-[15%] h-screen flex flex-col justify-start items-center">
            <div className={path == '/attendance' ? "bg-p-darkgray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"
            : "bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"}
            onClick={(e) => {router.push('/attendance')}}>
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Attendance</h1>
                </div>
            </div>
            <div className={path == '/leave-request' ? "bg-p-darkgray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"
            : "bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"}
            onClick={(e) => {router.push('/leave-request')}}>
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Leave Request</h1>
                </div>
            </div>
            <div className={path == '/attendance-history' ? "bg-p-darkgray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"
            : "bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"}
            onClick={(e) => {router.push('/attendance-history')}}>
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Attendance History</h1>
                </div>
            </div>
            <div className={path == '/leave-history' ? "bg-p-darkgray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"
            : "bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"}
            onClick={(e) => {router.push('/leave-history')}}>
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Leave History</h1>
                </div>
            </div>
            <div className={path == '/leave-management' ? "bg-p-darkgray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"
            : "bg-p-gray hover:bg-p-darkgray w-full h-[10%] text-center border-p-lightgray border-b-2 flex flex-row justify-center"}
            onClick={(e) => {router.push('/leave-management')}}>
                <div className="flex flex-col justify-center">
                    <h1 className="text-white text-lg">Leave Management</h1>
                </div>
            </div>
        </div>
    );
}