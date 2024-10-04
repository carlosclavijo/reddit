import React, { useEffect, useState } from "react"

function App() { 

  const url = "http://localhost:8080/users"
  const [users, setUsers] = useState([])

  const fetchData = () => {
    fetch(url)
      .then(response => {
        return response.json()
      })
      .then(data => {
        setUsers(data)
      })
  }

  useEffect(() => {
    fetchData()
  }, [])

  return (
    <div>
      {users.length > 0 && (
        <ul>
          {users.map(user => (
            <div>
              <li key={user.UserId}>Id: {user.UserId}</li>
              <li key={user.UserId}>Name: {user.Username}</li>
              <li key={user.UserId}>Email: {user.Email}</li>
              <li key={user.UserId}>PostKarma: {user.PostKarma}</li>
              <li key={user.UserId}>CommentKarma: {user.CommentKarma}</li>
              <li key={user.UserId}>AccountAvailable: {user.AccountAvailable}</li>
              <li key={user.UserId}>Admin: {user.Username}</li>
              <li key={user.UserId}>CreatedAt: {user.CreatedAt}</li>
              <li key={user.UserId}>UpdatedAt: {user.UpdatedAt}</li>
            </div>
          ))}
        </ul>
      )}
    </div>
  )
}

export default App;