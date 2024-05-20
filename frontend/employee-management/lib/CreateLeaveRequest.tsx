export default async function CreateLeaveRequest (token: string, eid: string, start: string, end: string, reasons: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_LEAVE_URL}/leaves`, {
        method: 'POST',
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            id: 2345,
            employee_id : eid,
            date_start: start,
            date_end: end,
            reason: reasons,
            status: "Pending"
        })
    });

    if (!response.ok) {
        throw new Error('Fuck you');
    }

    return await response.json();
}