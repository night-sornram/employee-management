export default async function GetLeaveByID (token: string, lid: number) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_LEAVE_URL}/api/leaves/${lid}`, {
        method: 'GET',
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error ('Cannot get leave, ran into problem');
    }
    return await response.json();
    
}