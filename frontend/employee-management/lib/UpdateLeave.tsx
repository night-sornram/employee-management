export default async function UpdateLeave (token: string, lid: number, status: string, opinion: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_LEAVE_URL}/api/leaves/${lid}`, {
        method: 'PUT',
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            status: status,
            manager_opinion: opinion
        })
    });
    if (!response.ok) {
        throw new Error ('Cannot update leave, ran into problem');
    }
    return await response.json();
}