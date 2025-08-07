# Golang-Practice-Lab-By-Rahul


Suppose we have a situation where we want to fetch user details using a user ID. Now, this user could come from different places — like a Postgres database, or maybe from a mock in memory when we are writing tests.

Instead of hardcoding the source of the user in our code, we create an interface — let’s say UserRepository. This interface just says, “I want something that can give me a user when I give it an ID.” It doesn’t care how or from where — just that it can do this one job.

Now I can write two different implementations:
	•	One real implementation, like PostgresRepo, that fetches users from a real database.
	•	One fake implementation, like MockRepo, that returns a dummy user — useful for testing without a database.

Then, in my service layer — let’s say UserService — I accept this UserRepository interface. That means my service doesn’t care where the user is coming from. It just says, “Hey, give me something that knows how to get a user.”

In production, I pass PostgresRepo{} to the service, and in testing, I pass MockRepo{}. The code remains the same — no changes needed. That’s the power of interfaces.

So in short, instead of coding to a specific type like PostgresRepo, I’m coding to an expected behavior — get user by ID. This makes my code more flexible, cleaner, and easier to test.