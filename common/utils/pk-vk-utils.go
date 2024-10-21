package utils

import (
	"io"
	"os"

	"github.com/celer-network/goutils/log"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
)

func ReadFileAndReturnByteArray(extractedFilePath string) ([]byte, error) {
	file, err := os.Open(extractedFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func ReadPlonkProvingKey(filename string, pk plonk.ProvingKey) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = pk.UnsafeReadFrom(f)
	return err
}

func WritePlonkProvingKey(filename string, pk plonk.ProvingKey) {
	f, err := os.Create(filename)
	if err != nil {
		log.Error("pk writing open failed... ", err)
	}
	_, err = pk.WriteTo(f)
	if err != nil {
		log.Error("pk writing failed... ", err)
	}
}

func ReadPlonkVerifyingKey(filename string, vk plonk.VerifyingKey) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = vk.UnsafeReadFrom(f)
	return err
}

func WritePlonkVerifyingKey(filename string, vk plonk.VerifyingKey) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = vk.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}

func ReadWitness(filename string, witness witness.Witness) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = witness.ReadFrom(f)
	return err
}

func WriteWitness(filename string, witness witness.Witness) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Error("witness writing failed... ", err)
		return err
	}

	_, err = witness.WriteTo(f)
	if err != nil {
		log.Error("witness writing failed... ", err)
		return err
	}
	return nil
}

func ReadPlonkCcs(filename string, ccs constraint.ConstraintSystem) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = ccs.ReadFrom(f)
	return err
}

func WritePlonkCcs(filename string, css constraint.ConstraintSystem) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = css.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}

func WritePlonkVKSolidity(filename string, vk plonk.VerifyingKey) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	err = vk.ExportSolidity(f)

	if err != nil {
		return err
	}
	return nil
}

func ReadPlonkProofFromLocalFile(filename string, proof plonk.Proof) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = proof.ReadFrom(f)
	return err
}

func WritePlonkProofIntoLocalFile(proof plonk.Proof, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Error(err)
		return err
	}

	defer f.Close()

	_, err = proof.WriteRawTo(f)

	if err != nil {
		log.Error("Failed to export groth16 proof")
		return err
	}
	return nil
}

func ReadProvingKey(filename string, pk groth16.ProvingKey) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = pk.UnsafeReadFrom(f)
	return err
}

func WriteProvingKey(filename string, pk groth16.ProvingKey) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = pk.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}

func ReadVerifyingKey(filename string, vk groth16.VerifyingKey) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = vk.UnsafeReadFrom(f)
	return err
}

func WriteVerifyingKey(filename string, vk groth16.VerifyingKey) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Error("vk writing failed... ", err)
		return err
	}

	_, err = vk.WriteTo(f)
	if err != nil {
		log.Error("vk writing failed... ", err)
		return err
	}
	return nil
}

func ReadCcs(filename string, ccs constraint.ConstraintSystem) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = ccs.ReadFrom(f)
	return err
}

func WriteCcs(filename string, css constraint.ConstraintSystem) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = css.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}

func WriteVKSolidity(filename string, vk groth16.VerifyingKey) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	err = vk.ExportSolidity(f)

	if err != nil {
		return err
	}
	return nil
}

func ReadProofFromLocalFile(filename string, proof groth16.Proof) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = proof.ReadFrom(f)
	return err
}

func WriteProofIntoLocalFile(filename string, proof groth16.Proof) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Error(err)
		return err
	}

	defer f.Close()

	_, err = proof.WriteRawTo(f)

	if err != nil {
		log.Error("Failed to export groth16 proof")
		return err
	}
	return nil
}

func ReadPublicInputs(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func WritePublicInputs(filename string, publicInputs []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Error(err)
		return err
	}

	defer f.Close()

	_, err = f.Write(publicInputs)

	if err != nil {
		log.Error("Failed to export groth16 proof")
		return err
	}
	return nil
}
