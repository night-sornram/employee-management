export default async function GetLeaveAdmin ( token: string , query: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_LEAVE_URL}/api/leaves/${query}`, {
        method: 'GET',
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
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