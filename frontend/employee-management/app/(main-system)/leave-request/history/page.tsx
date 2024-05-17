import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { Card, CardContent, CardHeader, CardDescription, CardTitle } from "@/components/ui/card";
import { Table, TableCaption, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Leave, UserJson } from "@/interface";
import GetMyLeaves from "@/lib/GetMyLeaves";
import GetUserProfile from "@/lib/GetUserProfile";
import dayjs from "dayjs";
import { getServerSession } from "next-auth";

export default async function Page() {

    const session = await getServerSession(authOptions);
    if (!session) return null;
    const userProfile:UserJson = await GetUserProfile(session?.user.token);
    const data: Leave[] = await GetMyLeaves(userProfile.employee_id, session.user.token);
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
                <Table>
                    <TableHeader>
                        <TableRow>
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
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            data.map((leave) => 
                            <TableRow>
                                <TableCell>
                                    {dayjs(leave.date_start).format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_end).format('DD/MM/YYYY')}
                                </TableCell>
                                <TableCell>
                                    {dayjs(leave.date_end).diff(dayjs(leave.date_start), 'day')}
                                </TableCell>
                                {
                                    leave.status == "Approved" ? 
                                    <TableCell className="text-green-600">
                                        {leave.status}
                                    </TableCell> : 
                                    leave.status == "Denied" ?
                                    <TableCell className="text-red-600">
                                        {leave.status}
                                    </TableCell> :
                                    <TableCell className="text-blue-600">
                                        {leave.status}
                                    </TableCell>
                                }
                            </TableRow>)
                        }
                    </TableBody>
                </Table>
            </div>
        </main>
    )
}