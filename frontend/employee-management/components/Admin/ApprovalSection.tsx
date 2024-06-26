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
import { useState } from "react";
import { Button } from "../ui/button";

export default function ApprovalSection ({data} : {data: Leave}) {

    const {data: session} = useSession();
    const router = useRouter();
    if(!session) {
        return null;
    }
    const [opinion, setOpinion] = useState('');

    return (
        <div className="space-y-5">
            <Button 
            onClick={() => {
                router.push("/dashboard/approve-leave")
            }
            }
            className=" w-20 flex justify-center items-center">
                Back
            </Button>
            <h1 className="text-2xl font-bold">Manager Approval</h1>
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
                data.status === "Pending" ?
                <div className="w-full flex flex-col space-y-3">
                    <Label htmlFor="opinion">Manager's Opinion</Label>
                    <Textarea id="opinion" onChange={(e) => setOpinion(e.target.value)}/>
                </div> : 
                <div className="w-full flex flex-col space-y-3">
                <Label htmlFor="opinion">Manager's Opinion</Label>
                <Textarea id="opinion" disabled defaultValue={data.manager_opinion ? data.manager_opinion : "-"}/>
            </div>
            }
            
            {
                data.status === "Pending"  ?
                <div className="flex flex-row space-x-3 justify-between w-full">
                    <button className="w-1/2 text-center bg-green-500 rounded px-3 py-2 text-white hover:bg-green-600 hover:shadow-md" 
                    onClick={(e) => handleApprove(session.user.token, data.id, "Approved", data.date_start, data.date_end, data.employee_id, opinion)}>Approve</button>
                    <button className="w-1/2 text-center bg-red-500 rounded px-3 py-2 text-white hover:bg-red-600 hover:shadow-md" 
                    onClick={(e) => handleApprove(session.user.token, data.id, "Denied", data.date_start, data.date_end, data.employee_id, opinion)}>Deny</button>
                </div> : 
                ""
            }
            
        </div>
    );
}