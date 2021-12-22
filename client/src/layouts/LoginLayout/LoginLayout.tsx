import React, { SyntheticEvent, useState } from 'react'
import Container from 'react-bootstrap/Container'
import { useNavigate } from 'react-router-dom'

import './LoginLayout.css'

function LoginLayout() {

    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    let navigate = useNavigate()

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault()

        const response = await fetch('http://localhost:8000/api/login', {
            method: 'POST',
            headers: {
                'Access-Control-Allow-Methods': 'GET,PUT,POST,DELETE,PATCH,OPTIONS',
                'Access-Control-Allow-Origin': '*',
                'Access-Control-Request-Headers':'*' ,
                'Content-Type': 'Application/Json',
                },
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        });

        const content = await response.json();

        navigate('/home');
        
    }




    return (
        <Container onSubmit={submit}>
            <main className="form-signin">
                <form>
                    <img className="mb-4" src="./assets/img/locker.svg" width={72} height={57} />
                    <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
                    <div className="form-floating">
                        <input type="email" className="form-control" id="floatingInput" placeholder="name@example.com" 
                        onChange={e => setEmail(e.target.value)}/>
                        <label htmlFor="floatingInput">Email address</label>
                    </div>
                    <div className="form-floating">
                        <input type="password" className="form-control" id="floatingPassword" placeholder="Password" 
                        onChange={e => setPassword(e.target.value)}/>
                        <label htmlFor="floatingPassword">Password</label>
                    </div>
                    <div className="checkbox mb-3">
                        <label>
                            <input type="checkbox" defaultValue="remember-me" /> Remember me
                        </label>
                    </div>
                    <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
                </form>
            </main>


        </Container>
    )
}

export default LoginLayout
