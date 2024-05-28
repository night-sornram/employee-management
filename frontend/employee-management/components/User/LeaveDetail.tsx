'use client'

import { Leave } from "@/interface";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import dayjs from "dayjs";
import { Textarea } from "../ui/textarea";
import UpdateLeave from "@/lib/UpdateLeave";
import { useSession } from "next-auth/react";
import { revalidateTag } from "next/cache";
import { useRouter } from "next/navigation";
import handleApprove from "@/lib/HandleApproval";
import { Button } from "../ui/button";

export default function LeaveDetail ({data} : {data: Leave}) {

    const {data: session} = useSession();
    const router = useRouter();
    if(!session) {
        return null;
    }


    return (
        <div className="space-y-3">
            <div className="w-full flex flex-col space-y-3">
                    <Label htmlFor="status">Status</Label>
                    <Input id="status" disabled type="text" defaultValue={data.status}/>
            </div>
            <div className="flex flex-row space-x-3 justify-between w-full">
                <div className="w-1/3 flex flex-col space-y-3">
                    <Label htmlFor="leave_start">Leave Start</Label>
                    <Input id="leave_start" disabled type="text" defaultValue={dayjs(data.date_start).format('DD-MM-YYYY')}/>
                </div>
                <div className="w-1/3 flex flex-col space-y-3">
                    <Label htmlFor="leave_end">Leave End</Label>
                    <Input id="leave_end" disabled type="text" defaultValue={dayjs(data.date_end).format('DD-MM-YYYY')}/>
                </div>
                <div className="w-1/3 flex flex-col space-y-3">
                    <Label htmlFor="duration">Duration</Label>
                    <Input id="duration" disabled type="text" defaultValue={(dayjs(data.date_end).diff(dayjs(data.date_start), 'day') + 1).toString()}/>
                </div>
                
            </div>
            <div className="w-full flex flex-col space-y-3">
                <Label htmlFor="reason">Reason</Label>
                <Textarea id="reason" disabled defaultValue={data.reason}/>
            </div>
            {
                data.status != "Pending" ? 
                <div className="w-full flex flex-col space-y-3">
                    <Label htmlFor="opinion">Manager's Opinion</Label>
                    <Textarea id="opinion" disabled defaultValue={data.manager_opinion ? data.manager_opinion : "-"}/>
                </div> : null
            }
            <div className="flex flex-row w-full">
                <Button className="w-full justify-center" onClick={() => router.push('/leave-request/history')}>Back</Button>
            </div>
        </div>
    );
}