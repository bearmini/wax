#[no_mangle]
pub extern fn add(a: u32, b: u32) -> u32 {
  a + b
}

#[no_mangle]
pub extern fn sub(a: u32, b: u32) -> u32 {
  a - b
}

#[no_mangle]
pub extern fn mul(a: u32, b: u32) -> u32 {
  a * b
}

#[no_mangle]
pub extern fn div(a: u32, b: u32) -> u32 {
  a / b
}
