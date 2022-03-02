
fn main() {
    let num :i32 = "42".parse().expect("Not an num");
   
    let tup :(i32, f64, u8) = (500, 6.4, 1);
   
    let (x, y, z) = tup;
    let arr:[i32; 5] = [1,2,3,4,5];
    // let a=[3;5];
    // let a=[3,3,3,3,3];
    println!("{}, {}, {}",x,y,z);
    println!("{} {}",num, x);
    for a in arr {
        print!("{} ",a)
    }
    another_function(23,32);
    let five = five();

    if five < 2 {
        println!("True");
    } else {
        println!("Fasle");
    }

    for elem in arr.iter() {
        print!("{}",elem);
    }
    for number in (1..4).rev() {
        print!("{} ",number);
    }
}

fn another_function(x:i32,y :i32) {
    println!("This is function! {} {}",x, y)
}

fn five() -> i32{
    5
}