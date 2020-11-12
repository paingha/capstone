// Port of https://www.rabbitmq.com/tutorials/tutorial-one-python.html. Run this
// in one shell, and run the hello_world_publish example in another.
use amiquip::{Connection, QueueDeclareOptions, ConsumerOptions, Result};
use std::thread;
use std::env;

fn exportFile() {
    println!("exporting now")
}

fn run_connection(mut connection: Connection) -> Result<()> {
    let channel = connection.open_channel(None)?;

    // Declaring the queue outside the thread spawn will fail, as it cannot
    // be moved into the thread. Instead, wait to declare until inside the new thread.

    // Would fail:
    // let queue = channel.queue_declare("hello", QueueDeclareOptions::default())?;
    thread::spawn(move || -> Result<()> {
        // Instead, declare once the channel is moved into this thread.
        let queue = channel.queue_declare("export", QueueDeclareOptions::default())?;
        let consumer = queue.consume(ConsumerOptions::default())?;
        for message in consumer.receiver().iter() {
            // do something with message...
            exportFile();
        }
        Ok(())
    });

    // do something to keep the connection open; if we drop the connection here,
    // it will be closed, killing the channel that we just moved into a new thread
}

fn main() -> Result<()> {
    // Open connection.
    let rabbit_url = env::var("ENV_RABBITMQ_HOST");
    let mut connection = Connection::insecure_open(rabbit_url)?;
    run_connection(connection);
    connection.close();
}


