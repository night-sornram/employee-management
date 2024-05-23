import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { Card, CardContent, CardHeader, CardDescription, CardTitle } from "@/components/ui/card";
import { Table, TableCaption, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Attendance, UserJson } from "@/interface";
import GetMyAttendances from "@/lib/GetMyAttendances";
import GetUserProfile from "@/lib/GetUserProfile";
import dayjs from "dayjs";
import { getServerSession } from "next-auth";
import utc from "dayjs/plugin/utc";
dayjs.extend(utc);

export default async function Page() {

    const session = await getServerSession(authOptions);
    if (!session) return null;
    const userProfile:UserJson = await GetUserProfile(session?.user.token);
    const data: Attendance[] = await GetMyAttendances(userProfile.employee_id, session.user.token);

    return(
        <main className='py-[5%] px-[5%] md:px-[10%] h-[93vh] md:w-[80%] 2xl:w-[60%] flex flex-col gap-10'>
            <div>
                <h1 className="font-bold text-2xl">
                    History of Attendance
                </h1>
            </div>
            <div className="flex flex-row gap-10 md:overflow-y-hidden overflow-y-scroll">
                <Card className="w-[320px] ">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            All-time Attendance
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {data.length}
                    </CardContent>
                </Card>
                <Card className="w-[320px]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            All-time Absence
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        0
                    </CardContent>
                </Card>
            </div>
            <div className="">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>
                                Date
                            </TableHead>
                            <TableHead>
                                Check-in Time
                            </TableHead>
                            <TableHead>
                                Check-out Time
                            </TableHead>
                            <TableHead>
                                Duration (Hour)
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            data.sort(function(a,b){
                                return Number(new Date(a.date)) - Number(new Date(b.date));
                            }).map((att) => 
                            <TableRow key={att.id}>
                                <TableCell>
                                    {dayjs(att.date).local().format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {att.leave_id !== -1 ? "LEAVE" : 
                                        (
                                            dayjs(att.check_in).local().format('HH:mm:ss')
                                        )
                                    }
                                    
                                </TableCell>
                                <TableCell>
                                    {   att.leave_id !== -1? "LEAVE" :
                                        (
                                            dayjs(att.check_out).local().toString() === "Mon, 01 Jan 0001 00:00:00 GMT" ?
                                            "-" :
                                            (
                                                dayjs(att.check_out).local().format('HH:mm:ss')
                                            )
                                        )
                                    }
                                    
                                </TableCell>
                                <TableCell>
                                    {
                                        att.leave_id !== -1? "LEAVE" 
                                        : 
                                        (
                                            dayjs(att.check_out).local().toString() === "Mon, 01 Jan 0001 00:00:00 GMT" ?
                                            "-" :
                                            (
                                                (Math.round(dayjs(att.check_out).diff(dayjs(att.check_in), 'hour', true) * 100) / 100).toFixed(2) +  " Hrs" 
                                            )
                                        )
                                    } 
                                </TableCell>
                            </TableRow>)
                        }
                    </TableBody>
                </Table>
            </div>
        </main>
    )
}