export interface RagChunk {
  id: string;
  sourcePath: string;
  content: string;
  embedding: number[];
}

export interface SearchMatch {
  chunk: RagChunk;
  score: number;
}

/**
 * Lightweight Cognitive RAG Vector Indexer & Semantic Search Engine
 * Splits text into chunks, generates deterministic token-hash embeddings,
 * and performs cosine similarity search.
 */
export class RagIndexer {
  private chunks: RagChunk[] = [];
  private dimension: number = 32;

  constructor() {
    this.seedDefaultKnowledge();
  }

  private hashToken(token: string): number[] {
    const vec = new Array(this.dimension).fill(0);
    for (let i = 0; i < token.length; i++) {
      const code = token.charCodeAt(i);
      const idx = (code * 31 + i) % this.dimension;
      vec[idx] += 1;
    }
    // Normalize vector
    const mag = Math.sqrt(vec.reduce((sum, v) => sum + v * v, 0)) || 1;
    return vec.map(v => v / mag);
  }

  public generateEmbedding(text: string): number[] {
    const tokens = text.toLowerCase().split(/\s+/).filter(Boolean);
    const combined = new Array(this.dimension).fill(0);
    if (tokens.length === 0) return combined;

    for (const t of tokens) {
      const tVec = this.hashToken(t);
      for (let i = 0; i < this.dimension; i++) {
        combined[i] += tVec[i];
      }
    }
    const mag = Math.sqrt(combined.reduce((sum, v) => sum + v * v, 0)) || 1;
    return combined.map(v => v / mag);
  }

  public cosineSimilarity(a: number[], b: number[]): number {
    let dot = 0;
    for (let i = 0; i < a.length; i++) {
      dot += a[i] * b[i];
    }
    return dot;
  }

  public addDocument(sourcePath: string, content: string): void {
    const lines = content.split("\n\n").filter(c => c.trim().length > 0);
    lines.forEach((part, idx) => {
      this.chunks.push({
        id: `${sourcePath}_${idx}`,
        sourcePath,
        content: part.trim(),
        embedding: this.generateEmbedding(part)
      });
    });
  }

  public search(query: string, topK: number = 3): SearchMatch[] {
    const qVec = this.generateEmbedding(query);
    const scored = this.chunks.map(chunk => ({
      chunk,
      score: this.cosineSimilarity(qVec, chunk.embedding)
    }));

    return scored
      .sort((a, b) => b.score - a.score)
      .slice(0, topK);
  }

  public getChunkCount(): number {
    return this.chunks.length;
  }

  private seedDefaultKnowledge(): void {
    this.addDocument("VISION.md", "Transforming the project into the ULTIMATE AGENTIC CODING HARNESS with 5-language parity.");
    this.addDocument("MEMORY.md", "Unified single-page WebUI dashboard consolidates all high-value features with tooltips.");
    this.addDocument("DEPLOY.md", "Nondestructive automated build and parity verification via scripts/build_and_verify_all.sh.");
  }
}
