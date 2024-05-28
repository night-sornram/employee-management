"use client"
import { Button } from "@/components/ui/button";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
  } from "@/components/ui/select"
import DatePickerRange from "@/components/Shadcnui/DatePickerRange";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { UserJson } from "@/interface";
import GetUserProfile from "@/lib/GetUserProfile";
import { useToast } from "@/components/ui/use-toast";
import { CreateLeaveRequestAction } from "../../../../lib/CreateRequestAction";
import CreateNotification from "@/lib/CreateNotification";
dayjs.extend(utc);

export default function LeaveRequestPage () {

    const current = dayjs().format('DD/MM/YYYY');

    const [start, setStart] = useState<Date>(new Date());
    const [end, setEnd] = useState<Date>(new Date());
    const [reason, setReason] = useState<string>('');
    const [category, setCategory] = useState<string>('');
    const router = useRouter();
    const {data: session} = useSession();
    const {toast} = useToast();
    if (!session) {
        router.push('/');
        return null;
    }    

    const submitHandler = async () => {
        if (reason == '') {
            return toast({
                variant: 'destructive',
                description: 'Please provide reason'
            });
        }
        const userProfile:UserJson = await GetUserProfile(session?.user.token);
        let startFormatted = dayjs(start).format('YYYY-MM-DDT[00:00:00Z]');
        let endFormatted = dayjs(end).format('YYYY-MM-DDT[00:00:00Z]');
        await CreateLeaveRequestAction(session.user.token, userProfile.employee_id, startFormatted, endFormatted, reason, category);
        await CreateNotification(session.user.token, "Leave Request", "You have new leave request", false, userProfile.employee_id);
        toast({
            title: "Request Submitted",
            description: "Your leave request has been submitted successfully",
          })
        setTimeout(() => {
            router.push('/leave-request/history');
        }
        , 1000)
    }

    return (
        <div className="flex flex-col space-y-5 px-[10%] py-[5%] md:w-[80%] 2xl:w-[60%] gap-[5%]">
            <h1 className="text-2xl font-bold">Leave Request</h1>
            <div className="flex flex-col space-y-3 justify-between w-full">
                <Label htmlFor="reason">Leave time</Label>
                <DatePickerRange className="w-full" onStartDateChange={setStart} onEndDateChange={setEnd}/>
            </div>
            <div className=" flex flex-col space-y-3">
                <Label htmlFor="reason">Category of Leave</Label>
                <Select onValueChange={(e) => setCategory(e.valueOf())}>
                    <SelectTrigger className="w-full">
                        <SelectValue placeholder="Category" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectItem value="Annual">Annual leave</SelectItem>
                        <SelectItem value="Maternity">Maternity leave</SelectItem>
                        <SelectItem value="Bereavement">Bereavement</SelectItem>
                        <SelectItem value="Sick">Sick leave</SelectItem>
                        <SelectItem value="Parental">Parental leave</SelectItem>
                        <SelectItem value="Unpaid">Unpaid leave</SelectItem>
                    </SelectContent>
                </Select>

            </div>
            <div>
                <Label htmlFor="reason">Reason of Leave</Label>
                <Textarea placeholder="Provide your reason here ..." id="reason" required={true} className="mt-3" onChange={(e) => setReason(e.target.value)}/>
            </div>
            <div className="items-center w-full text-center ">
                <Button className=" w-full flex justify-center" onClick={submitHandler}>Submit</Button>
            </div>
        </div>
    );
}