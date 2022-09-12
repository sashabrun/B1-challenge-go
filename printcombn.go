package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	for a := 48; a < 58; a++ {
		if n == 1 {
			z01.PrintRune(rune(a))
			if a != 57 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		} else {
			for b := 48; b < 58; b++ {
				if n == 2 {
					if a < b {
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						if a != 56 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				} else {
					for c := 48; c < 58; c++ {
						if n == 3 {
							if a < b && b < c {
								z01.PrintRune(rune(a))
								z01.PrintRune(rune(b))
								z01.PrintRune(rune(c))
								if a != 55 {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						} else {
							for d := 48; d < 58; d++ {
								if n == 4 {
									if a < b && b < c && c < d {
										z01.PrintRune(rune(a))
										z01.PrintRune(rune(b))
										z01.PrintRune(rune(c))
										z01.PrintRune(rune(d))
										if a != 54 {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								} else {
									for e := 48; e < 58; e++ {
										if n == 5 {
											if a < b && b < c && c < d && d < e {
												z01.PrintRune(rune(a))
												z01.PrintRune(rune(b))
												z01.PrintRune(rune(c))
												z01.PrintRune(rune(d))
												z01.PrintRune(rune(e))
												if a != 53 {
													z01.PrintRune(',')
													z01.PrintRune(' ')
												}
											}
										} else {
											for f := 48; f < 58; f++ {
												if n == 6 {
													if a < b && b < c && c < d && d < e && e < f {
														z01.PrintRune(rune(a))
														z01.PrintRune(rune(b))
														z01.PrintRune(rune(c))
														z01.PrintRune(rune(d))
														z01.PrintRune(rune(e))
														z01.PrintRune(rune(f))
														if a != 52 {
															z01.PrintRune(',')
															z01.PrintRune(' ')
														}
													}
												} else {
													for g := 48; g < 58; g++ {
														if n == 7 {
															if a < b && b < c && c < d && d < e && e < f && f < g {
																z01.PrintRune(rune(a))
																z01.PrintRune(rune(b))
																z01.PrintRune(rune(c))
																z01.PrintRune(rune(d))
																z01.PrintRune(rune(e))
																z01.PrintRune(rune(f))
																z01.PrintRune(rune(g))
																if a != 51 {
																	z01.PrintRune(',')
																	z01.PrintRune(' ')
																}
															}
														} else {
															for h := 48; h < 58; h++ {
																if n == 8 {
																	if a < b && b < c && c < d && d < e && e < f && f < g && g < h {
																		z01.PrintRune(rune(a))
																		z01.PrintRune(rune(b))
																		z01.PrintRune(rune(c))
																		z01.PrintRune(rune(d))
																		z01.PrintRune(rune(e))
																		z01.PrintRune(rune(f))
																		z01.PrintRune(rune(g))
																		z01.PrintRune(rune(h))
																		if a != 50 {
																			z01.PrintRune(',')
																			z01.PrintRune(' ')
																		}
																	}
																} else {
																	for i := 48; i < 58; i++ {
																		if n == 9 {
																			if a < b && b < c && c < d && d < e && e < f && f < g && g < h && h < i {
																				z01.PrintRune(rune(a))
																				z01.PrintRune(rune(b))
																				z01.PrintRune(rune(c))
																				z01.PrintRune(rune(d))
																				z01.PrintRune(rune(e))
																				z01.PrintRune(rune(f))
																				z01.PrintRune(rune(g))
																				z01.PrintRune(rune(h))
																				z01.PrintRune(rune(i))
																				if a != 49 {
																					z01.PrintRune(',')
																					z01.PrintRune(' ')
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
	z01.PrintRune('\n')
}
