export default async function GetMyLeaves (eid: string, token: string,query: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_LEAVE_URL}/api/leaves/me/${eid}${query}`, {
        method: 'GET',
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        next: {
            tags: ['leaves']
        }
    });
    if (!response.ok) {
        throw new Error("Cannot get your leaves for some reason");
    }
    return await response.json();
}