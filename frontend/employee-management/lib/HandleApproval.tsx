'use server'

import { revalidateTag } from "next/cache";
import UpdateLeave from "./UpdateLeave"
import { redirect } from "next/navigation";
import dayjs from "dayjs";
import CreateAttendance from "./CreateAttendance";

const handleApprove = async (token: string, lid: number, status: string, start: string, end: string, eid: string) => {
    await UpdateLeave(token, lid, status);

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
    }
    revalidateTag('leaves');
    redirect('/dashboard/approve-leave');
}

export default handleApprove;