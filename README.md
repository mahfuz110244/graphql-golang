# graphql-golang

## Generate 
```
go run github.com/99designs/gqlgen generate
```

## Playground
```
mutation createTodo {
  createTodo(input: {text: "todo", userId: "2"}) {
    user {
      id
    }
    text
    done
  }
}

query findTodos {
  todos {
    text
    done
    user {
      name
    }
  }
}

mutation createBook {
  createBook(
    input: {title: "Book5", price: 1000, isbn_no: "BK676757575869", author: {
      id: "2",
      name: "John",
      biography: "very good writer"
    }
  }
  ) {
    id
    title
    price
    isbn_no
    authors {
      id
      name
      biography
    }
  }
}

query GetAllBooksWithAuthors {
  books {
    id
    title
    price
    isbn_no
    authors {
      id
      name
      biography
    }
  }
}

query GetAllTheBooksOfAuthor1 {
authors(name: "author1") {
    books {
      id
      title
      price
      isbn_no
    }
	}
}

query GetAllTheBooksOfJohn {
authors(name: "John") {
    books {
      id
      title
      price
      isbn_no
    }
	}
}
```