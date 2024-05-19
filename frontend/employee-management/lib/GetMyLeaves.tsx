export default async function GetMyLeaves (eid: string, token: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_LEAVE_URL}/leaves/me/${eid}`, {
        method: 'GET',
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Cannot get your leaves for some reason");
    }
    return await response.json();
}