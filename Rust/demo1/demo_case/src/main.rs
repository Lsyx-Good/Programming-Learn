/*
 * @Author: Lsyx-Good hexiansi444@gmail.com
 * @Date: 2022-03-02 23:09:44
 * @LastEditors: Lsyx-Good hexiansi444@gmail.com
 * @LastEditTime: 2022-05-28 11:33:06
 * @FilePath: \demo1\demo_case\src\main.rs
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
fn main() {
    // s 不可用
    let  s = "Hello"; //s 可用
    // 可以对 s 进行相关操作 
    println!("{}",s);
    
    let mut s1 = String::from("Hello");

    s1.push_str(", World");
   
    println!("{}",s1);


} // s 作用域结束 s 不可用
