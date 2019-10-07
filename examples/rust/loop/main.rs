#[no_mangle]
pub extern fn infinite_loop() {
  let mut x = 0;
  loop {
    x += 1;
  }
}
