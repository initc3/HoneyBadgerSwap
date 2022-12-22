extern crate rand;
extern crate curve25519_dalek_ng;
use curve25519_dalek_ng::scalar::Scalar;
use curve25519_dalek_ng::ristretto::CompressedRistretto;
extern crate merlin;
use merlin::Transcript;
extern crate bulletproofs;
use bulletproofs::{BulletproofGens, PedersenGens, RangeProof};
use pyo3::prelude::*;
use curve25519_dalek_ng::ristretto::RistrettoPoint;
use curve25519_dalek_ng::traits::MultiscalarMul;

/// Given `data` with `len >= 32`, return the first 32 bytes.
pub fn read32(data: &[u8]) -> [u8; 32] {
    let mut buf32 = [0u8; 32];
    buf32[..].copy_from_slice(&data[..32]);
    buf32
}

// What will happen if we have more than 2^64 servers? Although it is extremely impossible, can
// adversarry exploit this?
pub fn lagrange_coefficient(covariate_u64: u64, others_u64: Vec<u64>) -> PyResult<Scalar> {
    let covariate = Scalar::from(covariate_u64);
    let others: Vec<_> = others_u64
        .iter()
        .map(|v| Scalar::from(*v))
        .collect();

    let denominator_factors: Vec<_> = others
        .iter()
        .map(|v| covariate - v)
        .collect();
    let numerator_factors: Vec<_> = others
        .iter()
        .map(|v| -v)
        .collect();
    
    let denominator = denominator_factors.iter().copied().reduce(|a, b| a * b).unwrap();
    let numerator = numerator_factors.iter().copied().reduce(|a, b| a * b).unwrap();

    let coefficient = numerator * denominator.invert();
    
    Ok(coefficient)
}

#[pyfunction]
fn pedersen_commit(secret_value_bytes: [u8; 32], blinding_bytes: [u8; 32]) -> PyResult<[u8; 32]> {
    let secret_value = Scalar::from_bytes_mod_order(secret_value_bytes);
    let blinding = Scalar::from_bytes_mod_order(blinding_bytes);

    let pc_gens = PedersenGens::default();
    let commitment = pc_gens.commit(secret_value, blinding);

    Ok(commitment.compress().to_bytes())
}

/// Provides an iterator over the powers of a `Scalar`.
///
/// This struct is created by the `exp_iter` function.
// pub struct ScalarExp {
//     x: Scalar,
//     next_exp_x: Scalar,
// }

// impl Iterator for ScalarExp {
//     type Item = Scalar;

//     fn next(&mut self) -> Option<Scalar> {
//         let exp_x = self.next_exp_x;
//         self.next_exp_x *= self.x;
//         Some(exp_x)
//     }

//     fn size_hint(&self) -> (usize, Option<usize>) {
//         (usize::max_value(), None)
//     }
// }

// /// Return an iterator of the powers of `x`.
// pub fn exp_iter(x: Scalar) -> ScalarExp {
//     let next_exp_x = Scalar::one();
//     ScalarExp { x, next_exp_x }
// }
/// Raises `x` to the power `n` using binary exponentiation,
/// with (1 to 2)*lg(n) scalar multiplications.
/// TODO: a consttime version of this would be awfully similar to a Montgomery ladder.
// pub fn scalar_exp_vartime(x: &Scalar, mut n: Scalar) -> Scalar {
//     let mut result = Scalar::one();
//     let mut aux = *x; // x, x^2, x^4, x^8, ...
//     let zer = Scalar::zero();
//     let on = Scalar::one();
//     let inv_2 = (Scalar::from(2u64)).invert().reduce();
//     while n > zer {
//         let bit = n & on;
//         if bit == on {
//             result = result * aux;
//         }
//         n = n * inv_2;
//         aux = aux * aux; // FIXME: one unnecessary mult at the last step here!
//     }
//     result
// }

        // z^0 * \vec(2)^n || z^1 * \vec(2)^n || ... || z^(m-1) * \vec(2)^n
        // let powers_of_2: Vec<Scalar> = util::exp_iter(Scalar::from(2u64)).take(n).collect();
// #[pyfunction]
// fn other_base_commit(cur_base_bytes: [u8; 32], secret_value_bytes: [u8; 32], blinding_bytes: [u8; 32]) -> PyResult<[u8; 32]> {
//     //cur_base ^ {secret_value_bytes} * h^{blinding_bytes}
//     let cur_base = Scalar::from_bytes_mod_order(cur_base_bytes);
//     let secret_value = Scalar::from_bytes_mod_order(secret_value_bytes);
//     let blinding = Scalar::from_bytes_mod_order(blinding_bytes);

//     let pow_of_cur_base: Scalar = scalar_exp_vartime(&cur_base, secret_value);
//     let zer = Scalar::zero();

//     let pc_gens = PedersenGens::default();
//     let blinding_com = pc_gens.commit(zer, blinding);

//     let commitment = pow_of_cur_base * blinding_com;

//     Ok(commitment.compress().to_bytes())
// }

#[pyfunction]
fn other_base_commit(g_x_bytes: [u8; 32], y_bytes: [u8; 32], blinding_bytes: [u8; 32]) -> PyResult<[u8; 32]> {
    // (g^x)^{secret_value} * h^{blinding} * g^{r}
    let g_x = CompressedRistretto(g_x_bytes).decompress().unwrap();
    let y = Scalar::from_bytes_mod_order(y_bytes);
    let blinding = Scalar::from_bytes_mod_order(blinding_bytes);

    let pc_gens = PedersenGens::default();

    let g_xy_h_rz_com = RistrettoPoint::multiscalar_mul(&[y, blinding], &[g_x, pc_gens.B_blinding]);


    Ok(g_xy_h_rz_com.compress().to_bytes())
}

#[pyfunction]
fn product_com(x_bytes: [u8; 32], y_bytes: [u8; 32]) -> PyResult<[u8; 32]> {
    // let x = CompressedRistretto(x_bytes).decompress().unwrap();
    // let y = CompressedRistretto(y_bytes).decompress().unwrap();
    let x = Scalar::from_bytes_mod_order(x_bytes);
    let y = Scalar::from_bytes_mod_order(y_bytes);

    let product_com = (x * y).reduce();
    
    Ok(product_com.to_bytes())
}


#[pyfunction]
fn pedersen_open(secret_value: u64, blinding: u64, commitment_bytes: [u8; 32]) -> PyResult<bool> {
    let secret_value_scalar = Scalar::from(secret_value);
    let blinding_scalar = Scalar::from(blinding);
    let pc_gens = PedersenGens::default();
    let opened_commitment_bytes = pc_gens.commit(secret_value_scalar, blinding_scalar).compress().to_bytes();
    Ok(opened_commitment_bytes == commitment_bytes)
}


#[pyfunction]
fn pedersen_aggregate(commitment_bytes_list: Vec<[u8; 32]>, covariates: Vec<u64>) -> PyResult<[u8;32]> {
    let commitment_list: Vec<_> = commitment_bytes_list
        .iter()
        .map(|v| CompressedRistretto(read32(v)).decompress().unwrap())
        .collect();

    let scaled_commitment_list: Vec<_> = commitment_list
        .iter()
        .enumerate()
        .map(|(idx, v)| v * lagrange_coefficient(covariates[idx], [&covariates[..idx], &covariates[idx+1..]].concat()).unwrap())
        .collect();

    let commitment = scaled_commitment_list.iter().copied().reduce(|a, b| a + b).unwrap();
    // TODO: figure out if the list is in the correct order.
    // for commitment_bytes in commitment_bytes_list.iter() {
    //     commitment_list.push(CompressedRistretto(read32(commitment_bytes)).decompress());
    // }

    // TODO: take in a list of Pedersen commitment. Interpolate them.
    // TODO: upgrade to robust version.
    Ok(commitment.compress().to_bytes())
}

#[pyfunction]
fn pedersen_compare(recovered_commitment: [u8; 32], onchain_commitment: [u8; 32]) -> PyResult<bool> {
    Ok(recovered_commitment == onchain_commitment)
}

// TODO: return the blinding as bytes using as_bytes (returns &[u8; 32])
#[pyfunction]
fn zkrp_prove(secret_value: u64, bits: usize) -> PyResult<(Vec<u8>, [u8; 32], [u8; 32])> {
    // Generators for Pedersen commitments.  These can be selected
    // independently of the Bulletproofs generators.
    let pc_gens = PedersenGens::default();

    // Generators for Bulletproofs, valid for proofs up to bitsize 64
    // and aggregation size up to 1.
    let bp_gens = BulletproofGens::new(64, 1);

    // The API takes a blinding factor for the commitment.
    let blinding = Scalar::random(&mut rand::thread_rng());

    // The proof can be chained to an existing transcript.
    // Here we create a transcript with a doctest domain separator.
    let mut prover_transcript = Transcript::new(b"zkrp");

    // Create a 32-bit rangeproof.
    let (proof, committed_value) = RangeProof::prove_single(
        &bp_gens,
        &pc_gens,
        &mut prover_transcript,
        secret_value,
        &blinding,
        bits,
    ).expect("A real program could handle errors");

    Ok((proof.to_bytes(), committed_value.to_bytes(), blinding.to_bytes()))
}

#[pyfunction]
fn zkrp_verify(proof_bytes: Vec<u8>, committed_value_bytes: [u8; 32]) -> PyResult<bool> {
    // Generators for Pedersen commitments.  These can be selected
    // independently of the Bulletproofs generators.
    let pc_gens = PedersenGens::default();

    // Generators for Bulletproofs, valid for proofs up to bitsize 64
    // and aggregation size up to 1.
    let bp_gens = BulletproofGens::new(64, 1);

	let proof = RangeProof::from_bytes(proof_bytes.as_slice()).expect("Error: Proof deserialization failed!");
    let committed_value = CompressedRistretto(read32(&committed_value_bytes));
    // Verification requires a transcript with identical initial state:
    let mut verifier_transcript = Transcript::new(b"zkrp");
    Ok(proof.verify_single(&bp_gens, &pc_gens, &mut verifier_transcript, &committed_value, 32).is_ok())
}

#[pyfunction]
fn gen_random_value(value_num: u64) -> PyResult<Vec<[u8; 32]>> {
    let mut res_val : Vec<[u8; 32]> = Vec::new();

    let mut i = 0;

    while i < value_num {
        let cur_res = Scalar::random(&mut rand::thread_rng());
        res_val.push(cur_res.to_bytes());
        i = i + 1;
    }

    Ok(res_val)
}


#[pyfunction]
fn zkrp_prove_mul(x_v: u64, y_v: u64, rx_prime_bytes: [u8; 32], ry_prime_bytes: [u8; 32]) -> PyResult<[u8; 32]> {
    let pc_gens = PedersenGens::default();

    let x = Scalar::from(x_v);
    let y = Scalar::from(y_v);
    let rx_prime = Scalar::from_bytes_mod_order(rx_prime_bytes);
    let ry_prime = Scalar::from_bytes_mod_order(ry_prime_bytes);

    let kx = Scalar::random(&mut rand::thread_rng());
    let kx_prime = Scalar::random(&mut rand::thread_rng());
    let ky = Scalar::random(&mut rand::thread_rng());
    let ky_prime = Scalar::random(&mut rand::thread_rng());
    let rz = Scalar::random(&mut rand::thread_rng());
    let kz_prime = Scalar::random(&mut rand::thread_rng());

    let com_kx = pc_gens.commit(kx, kx_prime).compress();
    let com_ky = pc_gens.commit(ky, ky_prime).compress();

    // com_kz = pc_gens.commit(x*ky, -rx_prime*ky + kz_prime);
    let t1_value:Scalar = (x * ky).reduce();
    let t1_blinding_value:Scalar = (kz_prime - (rx_prime * ky).reduce()).reduce();
    let com_kz = pc_gens.commit(t1_value, t1_blinding_value).compress();

    let mut prover_transcript = Transcript::new(b"zkrpmul");
    let mut c_bytes = [0u8; 64];
    prover_transcript.commit_bytes(b"com_kx",com_kx.as_bytes());
    prover_transcript.commit_bytes(b"com_ky",com_ky.as_bytes());
    prover_transcript.commit_bytes(b"com_kz",com_kz.as_bytes());
    prover_transcript.challenge_bytes(b"c", &mut c_bytes);
    let c = Scalar::from_bytes_mod_order_wide(&c_bytes);
    
    let sx:Scalar = ((c*x).reduce() + kx).reduce(); // sx = c * x + kx
    let sy:Scalar = ((c*y).reduce() + ky).reduce(); // sy = c * y + ky
    let sx_prime:Scalar = (kx_prime - (c*rx_prime).reduce()).reduce(); // sx' = c * (-rx') + kx'
    let sy_prime:Scalar = (ky_prime - (c*ry_prime).reduce()).reduce(); // sy' = c * (-ry') + ky'
    let sz_prime:Scalar = ((c*rz).reduce() + kz_prime).reduce(); // sz' = c * rz + kz'

    

    Ok(c.to_bytes())
}

#[pyfunction ]
fn zkrp_verify_mul(mx_bytes: [u8; 32], my_bytes: [u8; 32], c_rx_bytes: [u8; 32], c_ry_bytes: [u8; 32]) -> PyResult<()> {
    let mx = Scalar::from_bytes_mod_order(mx_bytes);
    let my = Scalar::from_bytes_mod_order(my_bytes);
    let c_rx = Scalar::from_bytes_mod_order(c_rx_bytes);
    let c_ry = Scalar::from_bytes_mod_order(c_ry_bytes);
    let zer = Scalar::zero();

    let pc_gens = PedersenGens::default();

    let g_mx = Scalar::from_bytes_mod_order(pc_gens.commit(mx,zer).compress().to_bytes());
    let c_x = (c_rx * (g_mx.invert())).reduce();
    let g_my = Scalar::from_bytes_mod_order(pc_gens.commit(my,zer).compress().to_bytes());
    let c_y = (c_ry * (g_my.invert())).reduce();

    Ok(())
}

/// A Python module implemented in Rust.
#[pymodule]
fn zkrp_pyo3(_py: Python, m: &PyModule) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(pedersen_commit, m)?)?;
    m.add_function(wrap_pyfunction!(pedersen_open, m)?)?;
    m.add_function(wrap_pyfunction!(pedersen_aggregate, m)?)?;
    m.add_function(wrap_pyfunction!(pedersen_compare, m)?)?;
    m.add_function(wrap_pyfunction!(zkrp_prove, m)?)?;
    m.add_function(wrap_pyfunction!(zkrp_verify, m)?)?;
    m.add_function(wrap_pyfunction!(gen_random_value, m)?)?;
    m.add_function(wrap_pyfunction!(zkrp_prove_mul, m)?)?;
    m.add_function(wrap_pyfunction!(zkrp_verify_mul, m)?)?;
    m.add_function(wrap_pyfunction!(other_base_commit, m)?)?;
    m.add_function(wrap_pyfunction!(product_com, m)?)?;
    Ok(())
}