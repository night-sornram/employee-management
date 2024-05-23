import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import ApprovalSection from "@/components/Admin/ApprovalSection";
import { Leave, UserJson } from "@/interface";
import GetLeaveByID from "@/lib/GetLeaveByID";
import GetUserProfile from "@/lib/GetUserProfile";
import { getServerSession } from "next-auth";
import { useSession } from "next-auth/react";
import { useEffect } from "react"

export default async function LeaveIDPage ({params} : {params: {lid: number}}) {

    const session = await getServerSession(authOptions);
    if (!session) {
        return null;
    }
    const data:Leave = await GetLeaveByID(session.user.token, params.lid);

    return (
        <main className="flex flex-col px-[10%] py-[5%] w-[60%] gap-[5%]">
            <h1 className="text-2xl font-bold">Manager Approval</h1>
            <ApprovalSection data={data}/>
        </main>
    );
}