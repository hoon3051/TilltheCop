import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import RouteComponent from './component/router'
import { BrowserRouter } from 'react-router-dom'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <RouteComponent />
    </BrowserRouter>
  </React.StrictMode>,
)
