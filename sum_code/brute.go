package main

import "fmt"

/*
Loads of optimisations can be done here.
Written very verbose to demonstrate the principals.
*/

func containsValue(m map[int]int, v int) bool {
	for _, x := range m {
		if x == v {
			return true
		}
	}
	return false
}

func main() {
	solve := make(map[int]int)

	for a := 1; a < 27; a++ {
		for clear := 0; clear < 26; clear++ {
			solve[clear] = 0
		}
		solve[0] = a
		for b := 1; b < 27; b++ {
			for clear := 1; clear < 26; clear++ {
				solve[clear] = 0
			}
			if containsValue(solve, b) {
				continue
			}
			solve[1] = b

			for c := 1; c < 27; c++ {
				for clear := 2; clear < 26; clear++ {
					solve[clear] = 0
				}
				if containsValue(solve, c) {
					continue
				}
				solve[2] = c

				if !(solve[0]*solve[1] == solve[2]) {
					continue
				}

				for d := 1; d < 27; d++ {
					for clear := 3; clear < 26; clear++ {
						solve[clear] = 0
					}
					if containsValue(solve, d) {
						continue
					}
					solve[3] = d
					for e := 1; e < 27; e++ {
						for clear := 4; clear < 26; clear++ {
							solve[clear] = 0
						}
						if containsValue(solve, e) {
							continue
						}
						solve[4] = e

						if !(solve[3]*solve[4] == solve[2]) {
							continue
						}

						for f := 1; f < 27; f++ {
							for clear := 5; clear < 26; clear++ {
								solve[clear] = 0
							}
							if containsValue(solve, f) {
								continue
							}
							solve[5] = f
							for g := 1; g < 27; g++ {
								for clear := 6; clear < 26; clear++ {
									solve[clear] = 0
								}
								if containsValue(solve, g) {
									continue
								}
								solve[6] = g

								if !(solve[5]*solve[6] == solve[2]) {
									continue
								}

								for h := 1; h < 27; h++ {
									for clear := 7; clear < 26; clear++ {
										solve[clear] = 0
									}
									if containsValue(solve, h) {
										continue
									}
									solve[7] = h
									for i := 1; i < 27; i++ {
										for clear := 8; clear < 26; clear++ {
											solve[clear] = 0
										}
										if containsValue(solve, i) {
											continue
										}
										solve[8] = i

										if !(solve[7]*solve[7] == solve[8]) {
											continue
										}

										for j := 1; j < 27; j++ {
											for clear := 9; clear < 26; clear++ {
												solve[clear] = 0
											}
											if containsValue(solve, j) {
												continue
											}
											solve[9] = j
											for k := 1; k < 27; k++ {
												for clear := 10; clear < 26; clear++ {
													solve[clear] = 0
												}
												if containsValue(solve, k) {
													continue
												}
												solve[10] = k

												if !(solve[8]+solve[9] == solve[10]) {
													continue
												}

												for l := 1; l < 27; l++ {
													for clear := 11; clear < 26; clear++ {
														solve[clear] = 0
													}
													if containsValue(solve, l) {
														continue
													}
													solve[11] = l

													if !(solve[11]*solve[6] == solve[10]) {
														continue
													}

													for m := 1; m < 27; m++ {
														for clear := 12; clear < 26; clear++ {
															solve[clear] = 0
														}
														if containsValue(solve, m) {
															continue
														}
														solve[12] = m

														if !(solve[0]*solve[3] == solve[12]) {
															continue
														}

														for n := 1; n < 27; n++ {
															for clear := 13; clear < 26; clear++ {
																solve[clear] = 0
															}
															if containsValue(solve, n) {
																continue
															}
															solve[13] = n
															for o := 1; o < 27; o++ {
																for clear := 14; clear < 26; clear++ {
																	solve[clear] = 0
																}
																if containsValue(solve, o) {
																	continue
																}
																solve[14] = o
																for p := 1; p < 27; p++ {
																	for clear := 15; clear < 26; clear++ {
																		solve[clear] = 0
																	}
																	if containsValue(solve, p) {
																		continue
																	}
																	solve[15] = p

																	if !(solve[14]*solve[3] == solve[15]) {
																		continue
																	}

																	for q := 1; q < 27; q++ {
																		for clear := 16; clear < 26; clear++ {
																			solve[clear] = 0
																		}
																		if containsValue(solve, q) {
																			continue
																		}
																		solve[16] = q
																		for r := 1; r < 27; r++ {
																			for clear := 17; clear < 26; clear++ {
																				solve[clear] = 0
																			}
																			if containsValue(solve, r) {
																				continue
																			}
																			solve[17] = r

																			if !(solve[6]*solve[16] == solve[17]) {
																				continue
																			}

																			for s := 1; s < 27; s++ {
																				for clear := 18; clear < 26; clear++ {
																					solve[clear] = 0
																				}
																				if containsValue(solve, s) {
																					continue
																				}
																				solve[18] = s
																				for t := 1; t < 27; t++ {
																					for clear := 19; clear < 26; clear++ {
																						solve[clear] = 0
																					}
																					if containsValue(solve, t) {
																						continue
																					}
																					solve[19] = t
																					for u := 1; u < 27; u++ {
																						for clear := 20; clear < 26; clear++ {
																							solve[clear] = 0
																						}
																						if containsValue(solve, u) {
																							continue
																						}
																						solve[20] = u

																						if !(solve[19]+solve[18] == solve[20]) {
																							continue
																						}

																						for v := 1; v < 27; v++ {
																							for clear := 21; clear < 26; clear++ {
																								solve[clear] = 0
																							}
																							if containsValue(solve, v) {
																								continue
																							}
																							solve[21] = v
																							for w := 1; w < 27; w++ {
																								for clear := 22; clear < 26; clear++ {
																									solve[clear] = 0
																								}
																								if containsValue(solve, w) {
																									continue
																								}
																								solve[22] = w

																								if !(solve[21]+solve[4] == solve[22]) {
																									continue
																								}

																								for x := 1; x < 27; x++ {
																									for clear := 23; clear < 26; clear++ {
																										solve[clear] = 0
																									}
																									if containsValue(solve, x) {
																										continue
																									}
																									solve[23] = x

																									if !(solve[18]+solve[3] == solve[23]) {
																										continue
																									}

																									for y := 1; y < 27; y++ {
																										for clear := 24; clear < 26; clear++ {
																											solve[clear] = 0
																										}
																										if containsValue(solve, y) {
																											continue
																										}
																										solve[24] = y
																										for z := 1; z < 27; z++ {
																											for clear := 25; clear < 26; clear++ {
																												solve[clear] = 0
																											}
																											if containsValue(solve, z) {
																												continue
																											}
																											solve[25] = z

																											if !((solve[24] + solve[22]) == (solve[17] + solve[25])) {
																												continue
																											}

																											fmt.Printf("%v\n", solve)
																										}

																									}

																								}

																							}

																						}

																					}

																				}

																			}

																		}

																	}

																}

															}

														}

													}

												}

											}

										}

									}

								}

							}

						}

					}

				}

			}

		}

	}
	fmt.Println("done")
}
