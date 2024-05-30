import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import ApprovalSection from "@/components/Admin/ApprovalSection";
import LeaveDetail from "@/components/User/LeaveDetail";
import { Leave, UserJson } from "@/interface";
import GetLeaveByID from "@/lib/GetLeaveByID";
import GetUserProfile from "@/lib/GetUserProfile";
import { getServerSession } from "next-auth";

export default async function LeaveIDPage ({params} : {params: {lid: number}}) {

    const session = await getServerSession(authOptions);
    if (!session) {
        return null;
    }
    const data:Leave = await GetLeaveByID(session.user.token, params.lid);

    return (
        <main className="py-[5%] px-[5%]  h-[93vh]  md:w-[70%] 2xl:w-[60%] flex flex-col gap-10">
            <ApprovalSection data={data}/>  
        </main>
    );
}