export default async function GetMyLeaves (eid: string, token: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_LEAVE_URL}/leaves/me/${eid}`, {
        method: 'GET',
        headers: {
            authorization: `Bearer ${token}`,
        },
    });
    if (!response.ok) {
        throw new Error("Cannot get your leaves for some reason");
    }
    return await response.json();
}