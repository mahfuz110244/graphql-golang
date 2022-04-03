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

type Task {
  id: ID!
  title: String!
  note: String!
  completed: Int!
  created_at: String!
  updated_at: String!
}

input NewTask {
  title: String!
  note: String!
}

type Author {
  id: ID!
  name: String!
  biography: String!
}

input NewAuthor {
  name: String!
  biography: String!
}

type Book {
  id: ID!
  title: String!
  price: Int!
  isbn_no: String!
  authors: Author!
}


input NewBook {
  title: String!
  price: Int!
  isbn_no: String!
  author_id: ID!
}

type Books {
  books: [Book]
}


type Mutation {
  createTask(input: NewTask!): Task!
  createAuthor(input: NewAuthor!): Author!
  createBook(input: NewBook!): Book!
}

type Query {
  tasks: [Task!]!
  books: [Book!]!
  authors(name : String!): Books!
}


CREATE TABLE authors (
	id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	"name" TEXT NOT NULL,
	biography TEXT
);

CREATE TABLE books (
	id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	title TEXT NOT NULL,
   price NUMERIC(10,2) NOT NULL,
	isbn_no TEXT NOT NULL,
   author_id INT NOT NULL,
   constraint fk_books_authors foreign key (author_id) references authors(id)
);
CREATE INDEX books_author_id_idx ON books(author_id);