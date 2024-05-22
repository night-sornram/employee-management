import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { Card, CardContent, CardHeader, CardDescription, CardTitle } from "@/components/ui/card";
import { Table, TableCaption, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Leave, UserJson } from "@/interface";
import GetLeaveAdmin from "@/lib/GetLeaveAdmin";
import GetUserProfile from "@/lib/GetUserProfile";
import dayjs from "dayjs";
import { getServerSession } from "next-auth";
import utc from "dayjs/plugin/utc";
import { LapTimerIcon , CheckIcon ,Cross1Icon  } from "@radix-ui/react-icons";
import Link from "next/link";
import { Button } from "@/components/ui/button";
dayjs.extend(utc);

export default async function Page() {

    const session = await getServerSession(authOptions);
    if (!session) return null;
    const userProfile:UserJson = await GetUserProfile(session?.user.token);
    const data: Leave[] = await GetLeaveAdmin(session.user.token);
    let approved = 0, denied = 0, pending = 0;
    for (let leave of data) {
        if (leave.status == 'Approved') {
            approved += 1;
        }
        if (leave.status == 'Denied') {
            denied += 1;
        }
        if (leave.status == 'Pending') {
            pending += 1;
        }
    }

    return(
        <main className=' p-10 h-[93vh] w-screen flex flex-col gap-10'>
            <div>
                <h1 className="font-bold text-2xl">
                    History of Leave
                </h1>
            </div>
            <div className="flex flex-row gap-10">
                <Card className="w-[20%]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Approved Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {approved}
                    </CardContent>
                </Card>
                <Card className="w-[20%]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Denied Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {denied}
                    </CardContent>
                </Card>
                <Card className="w-[20%]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Pending Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {pending}
                    </CardContent>
                </Card>
            </div>
            <div className="w-[60%]">
                <h1 className="font-bold text-2xl">
                    Pending Request
                </h1>
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>
                                Employee ID
                            </TableHead>
                            <TableHead>
                                From
                            </TableHead>
                            <TableHead>
                                To
                            </TableHead>
                            <TableHead>
                                Duration (days)
                            </TableHead>
                            <TableHead>
                                Status
                            </TableHead>
                            <TableHead>
                                Details
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            data.filter((leave) => leave.status == 'Pending').map((leave) =>
                            <TableRow  key={leave.id}>
                                <TableCell>
                                    {leave.employee_id}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_start).format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_end).format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_end).diff(dayjs(leave.date_start), 'day') + 1}
                                </TableCell>
                                {
                                    leave.status == "Approved" ? 
                                    <TableCell className=" flex flex-row">
                                        <CheckIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> : 
                                    leave.status == "Denied" ?
                                    <TableCell className=" flex flex-row">
                                        <Cross1Icon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> :
                                    <TableCell className=" flex flex-row">
                                        <LapTimerIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell>
                                }
                                <TableCell>                                    
                                    <Link href={`/dashboard/approve-leave/${leave.id}`} className="hover:underline text-sky-600">
                                        Details
                                    </Link>                                   
                                </TableCell>
                            </TableRow>)
                        }
                    </TableBody>
                </Table>
            </div>
            <div className="w-[60%]">
                <h1 className="font-bold text-2xl">
                    Success Request
                </h1>
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>
                                Employee ID
                            </TableHead>
                            <TableHead>
                                From
                            </TableHead>
                            <TableHead>
                                To
                            </TableHead>
                            <TableHead>
                                Duration (days)
                            </TableHead>
                            <TableHead>
                                Status
                            </TableHead>
                            <TableHead>
                                Details
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            data.filter((leave) => leave.status !== 'Pending').sort(
                                function(a,b){
                                    return Number(new Date(a.date_start)) - Number(new Date(b.date_start));
                                }
                            ).map((leave) => 
                            <TableRow key={leave.id}>
                                <TableCell>
                                    {leave.employee_id}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_start).format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_end).format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_end).diff(dayjs(leave.date_start), 'day') + 1}
                                </TableCell>
                                {
                                    leave.status == "Approved" ? 
                                    <TableCell className=" flex flex-row">
                                        <CheckIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> : 
                                    leave.status == "Denied" ?
                                    <TableCell className=" flex flex-row">
                                        <Cross1Icon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell> :
                                    <TableCell className=" flex flex-row">
                                        <LapTimerIcon className="mr-2 h-5 w-5"/> {leave.status}
                                    </TableCell>
                                }
                                <TableCell>                                    
                                    <Link href={`/dashboard/approve-leave/${leave.id}`} className="hover:underline text-sky-600">
                                        Details
                                    </Link>                                   
                                </TableCell>
                            </TableRow>)
                        }
                    </TableBody>
                </Table>
            </div>
        </main>
    )
}