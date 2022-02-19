fn main() {
    // s 不可用
    let  s = "Hello"; //s 可用
    // 可以对 s 进行相关操作 
    println!("{}",s);
    
    let mut s1 = String::from("Hello");

    s1.push_str(", World");
   
    println!("{}",s1);

       


} // s 作用域结束 s 不可用
