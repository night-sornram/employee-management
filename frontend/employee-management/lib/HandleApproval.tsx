'use server'

import { revalidateTag } from "next/cache";
import UpdateLeave from "./UpdateLeave"
import { redirect } from "next/navigation";
import dayjs from "dayjs";
import CreateAttendance from "./CreateAttendance";
import CreateNotification from "./CreateNotification";

const handleApprove = async (token: string, lid: number, status: string, start: string, end: string, eid: string, opinion: string) => {
    await UpdateLeave(token, lid, status, opinion);

    if (status == "Approved") {
        let startDate = dayjs(start);
        let endDate = dayjs(end);
        console.log(startDate.format());
        for (let date = startDate; date.isBefore(endDate) || date.isSame(endDate); date = date.add(1, 'day')) {
            let formattedDateTime = date.format('YYYY-MM-DDT[00:00:00Z]');
            let formattedDate = date.format('YYYY-MM-DD');
            console.log(formattedDateTime);
            console.log(formattedDate);
            await CreateAttendance(token, eid, formattedDateTime, formattedDate, lid);
        }
        await CreateNotification(token, "Leave Approved", `Your leave has been approved`, false, eid);

    }
    else {
        await CreateNotification(token, "Leave Denied", `Your leave has been denied`, false, eid);
    }
    revalidateTag('leaves');
    redirect('/dashboard/approve-leave');
}

export default handleApprove;