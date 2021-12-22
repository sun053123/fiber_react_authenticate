import React, {SyntheticEvent, useState} from 'react'
import Container from 'react-bootstrap/Container'
import { useNavigate } from 'react-router-dom'

function RegisterLayout() {

    const [username, setUsername] = useState<string>('')
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    let navigate = useNavigate();


    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        await fetch('http://localhost:8000/api/register', {
            method: 'POST',
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify({
                username,
                email,
                password,
            })
        });

        navigate('/login');
    }

    return (
        <Container >
            <main className="form-signin">
                <form onSubmit={submit}>
                    <img className="mb-4" src="./assets/img/locker.svg" width={72} height={57} />
                    <h1 className="h3 mb-3 fw-normal">Register </h1>
                    <div className="form-floating">
                        <input type="text" className="form-control" id="floatingInputUsername" placeholder="Username" 
                        onChange={e => setUsername(e.target.value)}/>
                        <label htmlFor="floatingInput">Username</label>
                    </div>
                    <div className="form-floating">
                        <input type="email" className="form-control" id="floatingInputEmail" placeholder="name@example.com" 
                        onChange={e => setEmail(e.target.value)}/>
                        <label htmlFor="floatingInput">Email address</label>
                    </div>
                    <div className="form-floating">
                        <input type="password" className="form-control" id="floatingPassword" placeholder="Password" 
                        onChange={e => setPassword(e.target.value)} />
                        <label htmlFor="floatingPassword">Password</label>
                    </div>
                    
                    <button className="w-100 btn btn-lg btn-primary" type="submit" >Register</button>
                </form>
            </main>


        </Container>
    )
}

export default RegisterLayout
