import { Card, CardContent, CardHeader, CardDescription, CardTitle } from "@/components/ui/card";
import { Table, TableCaption, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Leave } from "@/interface";

export default function Page() {

    const mockData: Leave[] = [{
        id: 1,
        employee_id: "E01",
        date_start: "13/05/2024",
        date_end: "15/05/2024",
        reason: "",
        status: "Approved",
        duration: 3
    },
    {
        id: 2,
        employee_id: "E01",
        date_start: "20/05/2024",
        date_end: "21/05/2024",
        reason: "",
        status: "Denied",
        duration: 2
    },
    {
        id: 3,
        employee_id: "E01",
        date_start: "28/05/2024",
        date_end: "30/05/2024",
        reason: "",
        status: "Pending",
        duration: 3
    }
]

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
                        1
                    </CardContent>
                </Card>
                <Card className="w-[20%]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Denied Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        1
                    </CardContent>
                </Card>
                <Card className="w-[20%]">
                    <CardHeader>
                        <CardTitle className="text-lg">
                            Pending Leave
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        1
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
                            mockData.map((att) => 
                            <TableRow>
                                <TableCell>
                                    {att.date_start}
                                </TableCell>
                                <TableCell>
                                    {att.date_end}
                                </TableCell>
                                <TableCell>
                                    {att.duration}
                                </TableCell>
                                {
                                    att.status == "Approved" ? 
                                    <TableCell className="text-green-600">
                                        {att.status}
                                    </TableCell> : 
                                    att.status == "Denied" ?
                                    <TableCell className="text-red-600">
                                        {att.status}
                                    </TableCell> :
                                    <TableCell className="text-blue-600">
                                        {att.status}
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