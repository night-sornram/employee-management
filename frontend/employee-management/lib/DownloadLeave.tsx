export default async function DownloadLeave (token: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_LEAVE_URL}/api/leaves/download`, {
        method: "GET",
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
        },
    });
    if (!response.ok) {
        throw new Error("Cannot get leaves for some reason");
    }
    // Read the response as a blob
    const blob = await response.blob();

    return blob;
}