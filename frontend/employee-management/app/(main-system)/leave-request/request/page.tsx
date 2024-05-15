"use client"
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import dayjs from "dayjs";

export default function LeaveRequestPage () {

    const current = dayjs().format('DD/MM/YYYY');

    return (
        <div className="flex flex-col px-[10%] py-[5%] w-[60%] gap-[10%]">
            <h1 className="text-2xl font-bold">Leave Request</h1>
            <div className="flex flex-row justify-between w-full">
                <div>
                    <Label htmlFor="from">From (first day of leave)</Label>
                    <Input type="date" id="from" className="mt-3" required={true}/>  
                </div>
                <div>
                    <Label htmlFor="to">To (last day of leave)</Label>
                    <Input type="date" id="to" className="mt-3" required={true}/>  
                </div>
            </div>
            <div>
                <Label htmlFor="reason">Reason of Leave</Label>
                <Textarea placeholder="Provide your reason here ..." id="reason" required={true} className="mt-3"/>
            </div>
            <div className="items-center w-full text-center ">
                <Button className=" w-full flex justify-center" >Submit</Button>
            </div>
        </div>
    );
}