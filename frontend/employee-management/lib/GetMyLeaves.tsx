export default async function GetMyLeaves (eid: string, token: string) {
    const response = await fetch(`http://127.0.0.1:8082/leaves/me/${eid}`, {
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