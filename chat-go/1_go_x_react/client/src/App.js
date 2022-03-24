import React from 'react'
import Navbar from './components/Navbar'
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import Item1 from './components/views/item1'
import Item2 from './components/views/item2'
import Item3 from './components/views/item3'
import Item4 from './components/views/item4'

import './app.css'

const Router = () => {
    return (
        <>
            <Switch>
                <Route exact path='/item1'>
                    <Item1 />
                </Route>
                <Route path='/item2'>
                    <Item2 />
                </Route>
                <Route path='/item3'>
                    <Item3 />
                </Route>
                <Route path='/item4'>
                    <Item4 />
                </Route>
            </Switch>
        </>
    )
}

export default function App() {
    return (
        <div>
            <BrowserRouter>
                <Navbar className='navbar'/>
                <hr/>
                <Router />
            </BrowserRouter>
        </div>
    )
}