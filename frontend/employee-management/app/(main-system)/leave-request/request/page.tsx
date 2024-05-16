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

export default function LeaveRequestPage () {

    const current = dayjs().format('DD/MM/YYYY');

    return (
        <div className="flex flex-col px-[10%] py-[5%] w-[60%] gap-[5%]">
            <h1 className="text-2xl font-bold">Leave Request</h1>
            <div className="flex flex-col space-y-3 justify-between w-full">
                <Label htmlFor="reason">Leave time</Label>
                <DatePickerRange className="w-full"/>
            </div>
            <div className=" flex flex-col space-y-3">
                <Label htmlFor="reason">Category of Leave</Label>
                <Select>
                    <SelectTrigger className="w-full">
                        <SelectValue placeholder="Category" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectItem value="annual_leave">Annual leave</SelectItem>
                        <SelectItem value="maternity_leave">Maternity leave</SelectItem>
                        <SelectItem value="bereavement">Bereavement</SelectItem>
                        <SelectItem value="sick_leave">Sick leave</SelectItem>
                        <SelectItem value="parental_leave">Parental leave</SelectItem>
                        <SelectItem value="unpaid_leave">Unpaid leave</SelectItem>
                    </SelectContent>
                </Select>

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