import { Card, CardContent, CardHeader, CardDescription, CardTitle } from "@/components/ui/card";
import { Table, TableCaption, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Attendance } from "@/interface";

export default function Page() {

    const mockData: Attendance[] = [{
        id: 1,
        employee_id: "E01",
        check_in: "8:30",
        check_out: "17:00",
        date: "19/05/2024",
        leave_id: null,
        duration: "8:21"
    },
    {
        id: 2,
        employee_id: "E02",
        check_in: "8:30",
        check_out: "17:00",
        date: "19/05/2024",
        leave_id: null,
        duration: "8:10"
    },
    {
        id: 3,
        employee_id: "E03",
        check_in: "8:30",
        check_out: "17:00",
        date: "19/05/2024",
        leave_id: null,
        duration: "8:04"
    }]

    return(
        <main className=' p-10 h-[93vh] w-screen flex flex-col gap-10'>
            <div>
                <h1 className="font-bold text-2xl">
                    History of Attendance
                </h1>
            </div>
            <div className="flex flex-row gap-10">
                <Card className="w-[20%]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            All-time Attendance
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {mockData.length}
                    </CardContent>
                </Card>
                <Card className="w-[20%]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            All-time Absence
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        {mockData.length - 3}
                    </CardContent>
                </Card>
            </div>
            <div className="w-[60%]">
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
                                Duration
                            </TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {
                            mockData.map((att) => 
                            <TableRow>
                                <TableCell>
                                    {att.date}
                                </TableCell>
                                <TableCell>
                                    {att.check_in}
                                </TableCell>
                                <TableCell>
                                    {att.check_out}
                                </TableCell>
                                <TableCell>
                                    {att.duration}
                                </TableCell>
                            </TableRow>)
                        }
                    </TableBody>
                </Table>
            </div>
        </main>
    )
}