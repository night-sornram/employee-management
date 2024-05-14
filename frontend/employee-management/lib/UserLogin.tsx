export default async function UserLogIn(userEmail : string, userPassword : string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ 
            email : userEmail,
            password : userPassword
        }),
    });

    if(!response.ok){
        throw new Error('Failed to login');
    }

    return await response.json();
}
