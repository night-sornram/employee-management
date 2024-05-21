'use server'

import { revalidateTag } from "next/cache";
import UpdateLeave from "./UpdateLeave"
import { redirect } from "next/navigation";

const handleApprove = async (token: string, lid: number, status: string) => {
    await UpdateLeave(token, lid, status);
    revalidateTag('leaves');
    redirect('/dashboard/approve-leave');
}

export default handleApprove;